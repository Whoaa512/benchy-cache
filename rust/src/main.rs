use std::fs;
use std::path::PathBuf;
use warp::{http::StatusCode, Filter};

const DATA_DIR: &'static str = "./data";

#[tokio::main]
async fn main() {
    let cache = warp::path("cache")
        .and(warp::path::param::<PathBuf>())
        .and(warp::fs::dir(DATA_DIR))
        .map(|_path: PathBuf, reply| warp::reply::with_status(reply, StatusCode::OK));

    let put_cache = warp::put()
        .and(warp::path("cache"))
        .and(warp::path::param::<String>())
        .and(warp::body::bytes())
        .map(|path: String, bytes: bytes::Bytes| {
            let file_path = format!("{}/{}", DATA_DIR, path);
            fs::write(file_path, bytes).expect("Unable to write file");
            warp::reply::with_status("File saved.", StatusCode::OK)
        });

    let status = warp::path("status").map(move || {
        let paths = fs::read_dir(DATA_DIR).expect("Unable to read directory");
        let mut total_size = 0u64;
        let mut item_count = 0u64;

        for path in paths {
            let meta = path
                .expect("Error reading path")
                .metadata()
                .expect("Error reading metadata");
            if meta.is_file() {
                total_size += meta.len();
                item_count += 1;
            }
        }

        warp::reply::json(&Status {
            size_on_disk: total_size,
            number_of_items: item_count,
        })
    });

    let routes = cache.or(put_cache).or(status);

    warp::serve(routes).run(([127, 0, 0, 1], 6942)).await;
}

#[derive(serde::Serialize)]
struct Status {
    size_on_disk: u64,
    number_of_items: u64,
}

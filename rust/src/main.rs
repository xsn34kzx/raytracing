use std::time::Instant;

fn main() {
    // Image
    let image_width = 256;
    let image_height = 256;

    // Render
    let render_start = Instant::now();

    println!("P3\n{} {}\n255", image_width, image_height);
    for j in 0..image_height {
        let v = j as f64 / (image_height - 1) as f64;
        for i in 0..image_width {
            let u = i as f64 / (image_width - 1) as f64;

            let r = (u * 255.0) as u8;
            let g = (v * 255.0) as u8;
            let b = 0;

            println!("{} {} {}", r, g, b);
        }
    }

    let render_end = render_start.elapsed();

    eprintln!("Finished in {} ms", render_end.as_millis());
}

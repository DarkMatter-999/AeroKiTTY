use sdl2::{event::Event, keyboard::Keycode};
use std::env;
use AeroKiTTY::{
    config::{FONT_SIZE, SCREEN_HEIGHT, SCREEN_WIDTH},
    frontend::Frontend,
    term::Term,
};

fn main() {
    // if cfg!(windows) {
    //     let shell = "C:\\Windows\\System32\\cmd.exe";
    //     run(shell.to_string());
    // } else if cfg!(unix) {
    //     let shell = env::var("SHELL").unwrap();
    //     run(shell);
    // }

    let shell = "cmd.exe".to_string();
    let mut term = Term::new(shell, "".to_string());

    let sdl_context = sdl2::init().unwrap();
    let video_subsys = sdl_context.video().unwrap();

    let window = video_subsys
        .window("AeroKiTTY", SCREEN_WIDTH, SCREEN_HEIGHT)
        .position_centered()
        .opengl()
        .build()
        .unwrap();

    let mut canvas = window
        .into_canvas()
        .build()
        .map_err(|e| e.to_string())
        .unwrap();

    let texture_creator = canvas.texture_creator();

    let ttf_context = sdl2::ttf::init().unwrap();
    let mut font = ttf_context
        .load_font("./fonts/font.ttf", FONT_SIZE)
        .unwrap();
    font.set_style(sdl2::ttf::FontStyle::NORMAL);

    canvas.clear();

    term.write_stdin();

    let mut frontend = Frontend::new(&mut canvas, &texture_creator, &font, 0, 0);

    'mainloop: loop {
        for event in sdl_context.event_pump().unwrap().poll_iter() {
            match event {
                Event::KeyDown {
                    keycode: Some(Keycode::Escape),
                    ..
                }
                | Event::Quit { .. } => break 'mainloop,
                _ => {}
            }
        }

        frontend.render(term.read_stdio());
    }

    println!("Exiting");
}

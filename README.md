# 🎮 Hide and Seek Countdown

A vibe-coded solution using [Cursor](https://www.cursor.com) (fun demo project for my kids). For real, Cursor did almost everything *including writing this README file* 😎

A countdown tool that screams numbers at your kids while they panic-hide! Features ASCII art, a progress bar, and a text-to-speech (TTS) voice that won't shut up until everyone's found their hiding spots.

## ✨ Features

- ⏲️ **The Unyielding Timekeeper (or, "No Cheating Allowed, Tiny Humans!")** - Tired of "9... 8... 7... DONE!"? This meticulously slow, perfectly enunciated countdown ensures even the most impatient or creatively-counting kiddos can't pull a fast one. You're welcome, parents.

- 🗣️ **The "Are You REALLY Ready?" Proclamation** - Before the seeker even *thinks* about opening their eyes, our program will loudly and dramatically announce, "Ready or not, here I come!" This gives the last-minute hiders a jolt and adds to the suspense. (May or may not include a unique voice for a dramatic curtain reveal.)

- ⏱️ **Adjustable "Patience" Settings** - Is your child a master hider who needs extra time, or a "pop-out-in-two-seconds" hider who needs less? Tailor the countdown speed from "Sloth Mode" (for the truly strategic) to "Lightning Round" (for the easily distracted). Simply specify the number of seconds you wish to count down from (set anywhere from 1–90 seconds).

- 🔊 **The Grand Announcer's Voice** - Because even a countdown deserves a dramatic reading!   Text-to-Speech (TTS) ensures everyone knows exactly how much time is left for prime hiding, even if they're stuffed under the couch.

- 📊 **The Suspense Meter** - Watch the tension build! Our visual progress bar goes down as the countdown ticks down, giving seekers a nail-biting, at-a-glance reminder of just how soon they become unleashed on the hiders.

- 🎯 **Secret Code Initiation** - No time for fumbling! With a specified Command-Line Argument, you can whisper the secret countdown time directly to the program for an instant game start. Perfect for impatient seekers (and even more impatient hiders!).

- 🔄 **"Again! Again!" Replay Button** - Why stop at one round of glorious hiding? Our Replay Functionality lets you jump straight into the next epic hide-and-seek battle with just a flick of the wrist. Endless fun (and endless opportunities to find new hiding spots)!

- 🌍 **Hide-and-Seek Everywhere!** - Whether you're hiding in a Mac, tucked away in a Windows PC, or camouflaged on a Linux machine, our program works seamlessly. No operating system can escape the joy of a good hide-and-seek game!

## 🚀 Quick Start

### Option 1: With Command-Line Argument
```bash
./hide_and_seek 15    # Start immediately with 15-second countdown
```

### Option 2: Interactive Mode
```bash
./hide_and_seek       # Prompts you to enter countdown time
```

## 📋 Requirements

- **Go 1.16+** (for building from source)
- **Text-to-Speech (Optional):**
  - **macOS**: Uses built-in `say` command (included)
  - **Windows**: Uses PowerShell's speech synthesis (included)
  - **Linux**: Requires `espeak` or `festival` package

### Installing Text-to-Speech on Linux
```bash
# Ubuntu/Debian
sudo apt-get install espeak

# Or alternatively
sudo apt-get install festival

# Fedora/CentOS
sudo yum install espeak
```

## 🛠️ Installation

### Build from Source
1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd hide_and_seek
   ```

2. **Build the application:**
   ```bash
   go build -o hide_and_seek main.go
   ```

3. **Run the application:**
   ```bash
   ./hide_and_seek
   ```

## 🎯 Usage Examples

### Quick 10-Second Game
```bash
./hide_and_seek 10
```

### Interactive Mode
```bash
./hide_and_seek
# Follow the prompts to enter your desired countdown time
```

### Multiple Games
The tool supports playing multiple rounds! After each game:
- Press `x` to exit
- Press `Enter` (or any other key) to play again
- **Note:** When replaying, you'll always be prompted for a new countdown time (interactive mode)

## 🎮 How to Play

1. **Start the Program**: Run `./hide_and_seek [optional-seconds]`
2. **Set Countdown Time**: Either via command-line argument or interactive prompt
3. **Get Ready**: Press Enter when everyone is ready to start hiding
4. **Hide!**: The countdown begins with visual and audio cues
5. **Seek!**: When the countdown reaches zero, it's time to find everyone!

## 🐛 Troubleshooting

### Text-to-Speech Not Working
- **macOS**: Should work automatically with built-in `say` command
- **Windows**: Requires PowerShell (usually pre-installed)
- **Linux**: Install `espeak` or `festival` package
- **Fallback**: Program continues in text-only mode if TTS fails

### Permission Issues
```bash
chmod +x hide_and_seek    # Make the binary executable
```

### Build Issues
- Ensure Go 1.16+ is installed (made for 1.21)
- Check that all dependencies are available
- Try cleaning and rebuilding: `go clean && go build -o hide_and_seek main.go`

## 📝 License

This project is licensed under the MIT License.

## 🤝 Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements!

## 🎉 Have Fun!

Enjoy your Hide and Seek games with your kids (or whomever)! 🎮✨

---
*Made with ❤️ for fun family game time!* 

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Global variable to track if TTS is available
var ttsAvailable bool = true

// speak uses the appropriate text-to-speech command for the current platform
func speak(text string) {
	if !ttsAvailable {
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("say", text)
	case "windows":
		// Use PowerShell's built-in speech synthesis
		psCommand := fmt.Sprintf("Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s')", text)
		cmd = exec.Command("powershell", "-c", psCommand)
	case "linux":
		// Try espeak first, fall back to festival if available
		if _, err := exec.LookPath("espeak"); err == nil {
			cmd = exec.Command("espeak", text)
		} else if _, err := exec.LookPath("festival"); err == nil {
			cmd = exec.Command("sh", "-c", fmt.Sprintf("echo '%s' | festival --tts", text))
		} else {
			ttsAvailable = false
			fmt.Println("🔇 Text-to-speech not available. Running in text-only mode.")
			return
		}
	default:
		ttsAvailable = false
		fmt.Println("🔇 Text-to-speech not supported on this platform. Running in text-only mode.")
		return
	}

	if err := cmd.Run(); err != nil {
		ttsAvailable = false
		fmt.Println("🔇 Text-to-speech failed. Running in text-only mode.")
	}
}

// speakWithVoice uses a specific voice for text-to-speech (macOS only)
func speakWithVoice(text string, voice string) {
	if !ttsAvailable || runtime.GOOS != "darwin" {
		// Fall back to regular speak function for non-macOS platforms
		speak(text)
		return
	}

	cmd := exec.Command("say", "-v", voice, text)
	if err := cmd.Run(); err != nil {
		// Fall back to regular speak if voice is not available
		speak(text)
	}
}

// clearScreen clears the terminal screen
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// printBanner prints a fun banner for the game
func printBanner() {
	banner := `
    ╔═════════════════════════════════════════════╗
    ║                                             ║
    ║   ██╗  ██╗██╗██████╗ ███████╗               ║
    ║   ██║  ██║██║██╔══██╗██╔════╝               ║
    ║   ███████║██║██║  ██║█████╗                 ║
    ║   ██╔══██║██║██║  ██║██╔══╝                 ║
    ║   ██║  ██║██║██████╔╝███████╗               ║
    ║   ╚═╝  ╚═╝╚═╝╚═════╝ ╚══════╝               ║
    ║                                             ║
    ║       █████╗ ███╗   ██╗██████╗              ║
    ║      ██╔══██╗████╗  ██║██╔══██╗             ║
    ║      ███████║██╔██╗ ██║██║  ██║             ║
    ║      ██╔══██║██║╚██╗██║██║  ██║             ║
    ║      ██║  ██║██║ ╚████║██████╔╝             ║
    ║      ╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝              ║
    ║                                             ║
    ║          ███████╗███████╗███████╗██╗  ██╗   ║
    ║          ██╔════╝██╔════╝██╔════╝██║ ██╔╝   ║
    ║          ███████╗█████╗  █████╗  █████╔╝    ║
    ║          ╚════██║██╔══╝  ██╔══╝  ██╔═██╗    ║
    ║          ███████║███████╗███████╗██║  ██╗   ║
    ║          ╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝   ║                             
    ║                                             ║
    ╚═════════════════════════════════════════════╝
`
	fmt.Println(banner)
}

// printCountdown displays the countdown with ASCII art
func printCountdown(seconds int) {
	// Create a box around the countdown number
	numberStr := fmt.Sprintf("%d", seconds)
	boxWidth := len(numberStr) + 8

	fmt.Println()
	fmt.Println("    " + strings.Repeat("═", boxWidth))
	fmt.Printf("    ║   %s   ║\n", numberStr)
	fmt.Println("    " + strings.Repeat("═", boxWidth))
	fmt.Println()
}

// printProgressBar shows a visual progress bar
func printProgressBar(current, total int) {
	barWidth := 40
	progress := float64(current) / float64(total)
	filled := int(float64(barWidth) * progress)

	fmt.Print("    Progress: [")
	fmt.Print(strings.Repeat("█", filled))
	fmt.Print(strings.Repeat("░", barWidth-filled))
	fmt.Printf("] %d%%\n", int(progress*100))
}

// validateCountdownTime validates the countdown time input
func validateCountdownTime(seconds int) error {
	if seconds < 1 || seconds > 90 {
		return fmt.Errorf("countdown time must be between 1 and 90 seconds")
	}
	return nil
}

// getCountdownTime prompts the user for countdown time with validation
func getCountdownTime() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("    ⏰ How many seconds should I count? (1-90): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("    ❌ Sorry, I couldn't read that. Please try again!")
			continue
		}

		// Clean the input
		input = strings.TrimSpace(input)

		// Convert to number
		seconds, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("    ❌ That's not a number! Please enter a number like 10 or 30.")
			continue
		}

		// Validate range
		if err := validateCountdownTime(seconds); err != nil {
			fmt.Printf("    ❌ %s\n", err.Error())
			continue
		}

		return seconds
	}
}

// countdown performs the actual countdown with visual and audio feedback
func countdown(seconds int) {
	// Initial announcement
	fmt.Println("\n    🎯 Starting countdown...")
	speak("OK! You better hide! I am counting!")

	// Give players time to start hiding
	time.Sleep(2 * time.Second)

	// Countdown loop
	for i := seconds; i > 0; i-- {
		clearScreen()
		printBanner()

		// Visual countdown display with ASCII art
		printCountdown(i)
		printProgressBar(i, seconds)

		// Add some visual flair based on remaining time
		if i <= 3 {
			fmt.Println("    ⚠️  ALMOST READY TO FIND YOU! ⚠️")
			fmt.Println("    🏃‍♂️  Better be hidden!")
		} else if i <= 10 {
			fmt.Println("    🏃 Better find a good hiding spot!")
		} else {
			fmt.Println("    😴 Take your time finding the perfect spot...")
		}

		// Speak the number
		speak(strconv.Itoa(i))

		// Wait between counts (slightly less than a second to account for speech time)
		time.Sleep(800 * time.Millisecond)
	}

	// Final announcement
	clearScreen()
	printBanner()
	fmt.Println("\n🚀 TIME'S UP! 🚀")
	fmt.Println("👀 Ready or not, here I come!")
	speakWithVoice("Ready or not, here I come!", "Trinoids")

	// Keep the final message visible
	time.Sleep(3 * time.Second)
}

// playGame runs one complete game session
func playGame(predefinedSeconds ...int) {
	var seconds int

	// If predefined seconds provided, use that; otherwise prompt user
	if len(predefinedSeconds) > 0 {
		seconds = predefinedSeconds[0]
		fmt.Printf("⏰ Counting down from %d seconds (from command line argument)\n", seconds)
	} else {
		// Get countdown time from user
		seconds = getCountdownTime()
	}

	// Confirm and start
	fmt.Printf("\n✅ Great! I'll count down from %d seconds.\n", seconds)
	fmt.Print("Press Enter (or the Return key) when everyone is ready to start hiding...")
	speak("Press Enter (or the Return key) when everyone is ready to start hiding")
	bufio.NewReader(os.Stdin).ReadString('\n')

	// Start countdown
	countdown(seconds)

	fmt.Println("\n🎉 Have fun playing Hide and Seek! 🎉")
}

func main() {
	// Clear screen and show banner
	clearScreen()
	printBanner()

	fmt.Println("Welcome to Hide and Seek Countdown! 🎮")
	fmt.Println("This tool will help you count down before seeking!")
	fmt.Println()

	// Check for command-line argument for countdown seconds
	var predefinedSeconds int
	var hasArgument bool

	if len(os.Args) > 1 {
		if seconds, err := strconv.Atoi(os.Args[1]); err == nil {
			if err := validateCountdownTime(seconds); err == nil {
				predefinedSeconds = seconds
				hasArgument = true
				fmt.Printf("🎯 Using %d seconds from command line argument\n\n", seconds)
			} else {
				fmt.Printf("❌ Invalid countdown time argument: %s\n", err.Error())
				fmt.Println("💡 Usage: ./hide_and_seek [seconds] (where seconds is between 1-90)")
				fmt.Println("Continuing with interactive mode...\n")
			}
		} else {
			fmt.Printf("❌ Invalid argument '%s' - must be a number\n", os.Args[1])
			fmt.Println("💡 Usage: ./hide_and_seek [seconds] (where seconds is between 1-90)")
			fmt.Println("Continuing with interactive mode...\n")
		}
	}

	// Game loop
	for {
		if hasArgument {
			playGame(predefinedSeconds)
			// After first game, always use interactive mode
			hasArgument = false
		} else {
			playGame()
		}

		// Ask if user wants to play again
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("🎮 Press 'x' to exit or hit Enter (or the Return key) to play again!")
		//speak("Press X to exit or any other key to play again")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "x" {
			fmt.Println("\n👋 Thanks for playing! Goodbye! 👋")
			speak("Thanks for playing! Goodbye!")
			break
		}

		// Clear screen for next game
		clearScreen()
		printBanner()
		fmt.Println("🎮 Starting a new game! 🎮")
		fmt.Println()
	}
}

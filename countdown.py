# This was my first draft which I wrote the old-school way (i.e., without Cursor).
# You can ignore this file. It was made specifically for macOS.

import subprocess
import time
import sys

def countdown(n):
    for i in range(n, 0, -1):
        subprocess.run(["say", str(i)])
        time.sleep(1)
    subprocess.run(["say", "-v", "Organ", "Ready or not, here I come!"])

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 countdown.py <number>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
        if n <= 0:
            raise ValueError("Number must be positive")
    except ValueError:
        print("Please provide a valid positive integer")
        sys.exit(1)
    
    countdown(n)

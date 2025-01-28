# References:
# https://medium.com/@noransaber685/simple-guide-to-creating-a-command-line-interface-cli-in-python-c2de7b8f5e05

import os
import cmd
import argparse
from datetime import datetime

class HelloWorldCLI(cmd.Cmd):

    prompt = "> "
    intro = "Welcome to the Hello World CLI! Type \"help\" for available commands."

    def __init__(self):
        super().__init__()
        self.current_directory = os.getcwd()

    def do_greeting(self, line):
        """Print a greeting!"""
        print("🌍 Hello world!")

    def do_get_current_directory(self, line):
        """Prints the current directory."""
        print(f"📂 {self.current_directory}")

    def do_get_current_time(self, line):
        """Prints the current time."""
        print(f"⏰ {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

    def do_quit(self, line):
        """Exit interactive CLI."""
        print("👋 Goodbye!")
        return True

    def emptyline(self):
        """Prevent running the last command when Enter is pressed with no input."""
        print("⚠️ Please enter a command or type 'help' for assistance.")

def parse_arguments():
    """Set up and parse command-line arguments."""
    parser = argparse.ArgumentParser(description="HelloWorldCLI - A simple command-line interface.")
    parser.add_argument("--greeting", action="store_true", help="Display a greeting message and exit.")
    parser.add_argument("--get-current-directory", action="store_true", help="Print the current working directory and exit.")
    parser.add_argument("--get-current-time", action="store_true", help="Print the current time and exit.")
    parser.add_argument("--interactive", action="store_true", help="Run the CLI interactively.")
    return parser.parse_args()

def main():
    args = parse_arguments()

    if args.greeting:
        HelloWorldCLI().do_greeting(None)
        return
    
    if args.get_current_directory:
        HelloWorldCLI().do_get_current_directory(None)
        return
    
    if args.get_current_time:
        HelloWorldCLI().do_get_current_time(None)
        return
    
    if args.interactive or not any(vars(args).values()):
        HelloWorldCLI().cmdloop()

if __name__ == "__main__":
    main()


class Dial:
    def __init__(self, start: int, dial_count: int):
        self.dial = start
        self.dial_count = dial_count
        self.zero_counter = 0
    
    def rotate(self, command: str):
        direction = command[0]
        amount = int(command[1:])
        
        if direction == 'L':
            amount = -amount
        
        # Position before wrapping. Can be far outside the valid range.
        new_raw = self.dial + amount

        if amount > 0:
            # Moving forward. Count how many full dial_size steps were passed:
            # check how many multiples of dial_size are between old and new_raw.
            crossings = (new_raw // self.dial_count) - (self.dial // self.dial_count)
        elif amount < 0:
            # Moving backward. Same idea, but with the reversed interval.
            crossings = ((self.dial - 1) // self.dial_count) - ((new_raw - 1) // self.dial_count)
        else:
            crossings = 0
        self.zero_counter += crossings

        # Wrap final value into the valid range.
        self.dial = new_raw % self.dial_count

    def get_zero_counter(self) -> int:
        return self.zero_counter


def main():
    dial = Dial(50, 100)
    with open('source.txt', 'r') as text_file:
        for row in text_file:
            command = row.strip()
            dial.rotate(command)
                
    print(f'The Password is: {dial.get_zero_counter()}')


if __name__ == "__main__":
    main()

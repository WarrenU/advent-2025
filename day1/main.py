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
        
        # rotate with wrap
        self.dial = (self.dial + amount) % self.dial_count
        if self.dial == 0:
            self.zero_counter += 1

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

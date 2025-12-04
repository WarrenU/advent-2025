def main():
    dial_val = 50
    zero_counter = 0
    with open('test.txt', 'r') as text_file:
        for line in text_file:
            row = line.strip()
            delim = row[0]
            rotations = int(row[1:])
            add = -rotations if delim == 'L' else rotations
            dial_val += add
            delim = 1 * int(str(dial_val)[0]) if dial_val > 100 else 1
            delim = 1 * int(str(dial_val)[1]) if dial_val < -99 else 1
            print(delim)
            if dial_val > 99:
                dial_val -= 100 * delim
            elif dial_val < 1:
                dial_val += 100 * delim
            if dial_val == 0 or dial_val == 100:
                dial_val = 0
                zero_counter += 1
            print(dial_val, zero_counter)
            
                
        print(f'The Password is: {zero_counter}')

if __name__ == "__main__":
    main()
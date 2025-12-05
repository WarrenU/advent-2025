def is_id_valid(id: int) -> bool:
    id_as_str = str(id)
    id_len = len(id_as_str)
    mid_index = int(id_len/2)
    for i in range(mid_index, 0, -1):
        potential_repeating_substring = id_as_str[:i]
        occurrences = id_as_str.count(potential_repeating_substring)
        chunk_count = id_len/i
        # print(f'CHUNK: {potential_repeating_substring} CHUNK_COUNT: {chunk_count}')
        if occurrences == chunk_count:
            return False
    return True

def main():
    invalid_ids = []
    with open('input.csv', 'r') as csv:
        for row in csv:
            ids = row.split(',')
            for id_range in ids:
                vals = id_range.split('-')
                start = int(vals[0])
                end = int(vals[1])
                for id in range(start, end+1):
                   if not is_id_valid(id):
                        invalid_ids.append(id)
        print(sum(invalid_ids))


if __name__ == "__main__":
    main()
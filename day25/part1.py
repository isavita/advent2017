
import re

def parse_input(file_path):
    with open(file_path, 'r') as file:
        lines = file.read().splitlines()
    initial_state = lines[0][-2]
    steps = int(re.findall(r'\d+', lines[1])[0])
    states = {}
    for i in range(3, len(lines), 10):
        state = lines[i][-2]
        value_0 = int(lines[i+2][-2])
        move_0 = -1 if lines[i+3].split()[-1] == 'left.' else 1
        next_state_0 = lines[i+4][-2]
        value_1 = int(lines[i+6][-2])
        move_1 = -1 if lines[i+7].split()[-1] == 'left.' else 1
        next_state_1 = lines[i+8][-2]
        states[state] = {(0,): (value_0, move_0, next_state_0), (1,): (value_1, move_1, next_state_1)}
    return initial_state, steps, states

def run_turing_machine(file_path):
    state, steps, states = parse_input(file_path)
    tape = {}
    cursor = 0
    for _ in range(steps):
        value = tape.get(cursor, 0)
        new_value, move, next_state = states[state][(value,)]
        tape[cursor] = new_value
        cursor += move
        state = next_state
    checksum = sum(tape.values())
    return checksum

print(run_turing_machine("day25/input.txt"))

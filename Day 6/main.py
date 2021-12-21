from copy import copy

def count_births_from_fish(days_until_next_birth, days, days_between_births=6):
    return int( (days + days_until_next_birth) / days_between_births )

def simulate_fish(fish, days, days_between_births=6, init_birth_wait=8):
    fish = copy(fish)
    n_births = 0

    while days > 0:
        # Births
        for i in range(len(fish)):
            if fish[i] == 0:
                n_births += 1
                fish[i] = days_between_births
            else:
                fish[i] -= 1

        for _ in range(n_births):
            fish += [init_birth_wait]

        n_births = 0
        days -= 1
    
    print(fish)
    return len(fish)

with open(file="test_input2.txt") as f:
    lines = f.readline()

numbers = [int(s) for s in 
    [e[0] for e in lines.split(",")]]

# Part 1
# for i in range(1,9):
#     n_fish = simulate_fish(numbers, i, 2, 3)
#     #print(f"{i}:\tn = {n_fish}")

# print(count_births_from_fish(1, 8, 3))

days = 5
a = 6
input = numbers

population = len(input)

queue = [(e,days) for e in input]
while len(queue) > 0:
    b, d = queue.pop(0)
    print((b, d)) 
    
    births = count_births_from_fish(b, d, a)
    population += births

    for i in range(1,births+1):
        if d - a*i > 0:
            queue += [(a, d-a*i)]

print(f"Population: {population}")
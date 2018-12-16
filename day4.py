import datetime

input_file = 'day4_input.txt'
# input_file = 'day4_input_test.txt'


class Guard(object):
    def __init__(self, guard_id):
        self.guard_id = guard_id
        self.minutes_asleep = 0
        self.sleep_by_minute = {}

    def __repr__(self):
        return 'Guard #%s, minutes asleep: %s\n%s' % (self.guard_id, self.minutes_asleep, self.sleep_by_minute)

    def populate_sleep_by_minute(self, start_time, duration):
        for minute in range(start_time, start_time+duration):
            try:
                self.sleep_by_minute[minute] += 1
            except KeyError:
                self.sleep_by_minute[minute] = 1


with open(input_file, 'r') as fd:
    action_log = fd.read().splitlines()

action_log.sort()
print('\n'.join(action_log))
guards = {}
current_guard = None
sleep_start = None
sleep_stop = None

for action in action_log:
    date_str = action.split()[0][1:]
    time_str = action.split()[1][:-1]
    log_type = action.split()[2].lower()
    guard_id = action.split()[3][1:]

    # switch guards
    if log_type == 'guard':
        try:
            current_guard = guards[guard_id]
        except KeyError:
            guards[guard_id] = Guard(guard_id)
            current_guard = guards[guard_id]
        finally:
            sleep_start = None
            sleep_stop = None
    # guard falls asleep
    elif log_type == 'falls':
        sleep_start = datetime.datetime.strptime('%s %s' % (date_str, time_str), '%Y-%m-%d %H:%M')
    # guard wakes up
    elif log_type == 'wakes':
        sleep_stop = datetime.datetime.strptime('%s %s' % (date_str, time_str), '%Y-%m-%d %H:%M')
        minutes_asleep = int((sleep_stop - sleep_start).total_seconds())//60
        current_guard.minutes_asleep += minutes_asleep
        current_guard.populate_sleep_by_minute(sleep_start.minute, minutes_asleep)

# find guard who has slept the most
max_sleep_time = 0
sleepiest_guard = None
for guard in guards.values():
    if guard.minutes_asleep > max_sleep_time:
        max_sleep_time = guard.minutes_asleep
        sleepiest_guard = guard


# find the sleepiest minute for the sleepiest guard
sleepiest_minute = None
for minute in sleepiest_guard.sleep_by_minute:
    # print(minute)
    if sleepiest_minute:
        if sleepiest_guard.sleep_by_minute[minute] > sleepiest_guard.sleep_by_minute[sleepiest_minute]:
            sleepiest_minute = minute
    else:
        sleepiest_minute = minute

print(sleepiest_guard)
print(sleepiest_minute)

# part two
sleepiest_guard = None
sleepiest_minute = None
sleepiest_minute_count = 0
for guard in guards.values():
    for minute in guard.sleep_by_minute:
        if guard.sleep_by_minute[minute] > sleepiest_minute_count:
            sleepiest_guard = guard
            sleepiest_minute = minute
            sleepiest_minute_count = sleepiest_guard.sleep_by_minute[sleepiest_minute]

print(sleepiest_guard)
print(sleepiest_minute)

require 'date'

EVENT_RE = /\[(?<time>[^\]]*)\] (Guard #(?<id>\d+) )?(?<event>.*)/

input = File.read('input.txt')
input = <<-INPUT
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
INPUT
raw_events = input.split("\n")

events = raw_events.map do |raw_event|
  match = EVENT_RE.match(raw_event)
  event = match.names.zip(match.captures).to_h
  event['time'] = DateTime.parse(event['time'])
  event['id'] = event['id'].to_i if event['id']
  event
end

events.sort_by! { |event| event['time'] }

guard_index = Hash.new { |h, k| h[k] = [] }
current_id = nil
events.each do |event|
  current_id = event['id'] if event['id']
  guard_index[current_id] << event
end

guard_times = Hash.new(0)
guard_index.each do |id, guard_events|
  guard_events.each_with_index do |event, i|
    next if ['begins shift', 'wakes up'].include?(event['event'])

    next_event = guard_events[i + 1]
    raise 'event order mismatch' if next_event['event'] != 'wakes up'

    guard_times[id] += next_event['time'].to_time - event['time'].to_time
  end
end

sleepiest_guard_id = guard_times.max_by(&:last).first
sleepiest_guard = guard_index[sleepiest_guard_id]

sleep_minutes = Hash.new(0)
sleepiest_guard.each_with_index do |event, i|
  next if ['begins shift', 'wakes up'].include?(event['event'])

  next_event = sleepiest_guard[i + 1]
  raise 'event order mismatch' if next_event['event'] != 'wakes up'

  seconds_asleep = next_event['time'].to_time - event['time'].to_time
  minutes_asleep = seconds_asleep.to_i / 60
  minutes_asleep.to_i.times do |minute|
    sleep_minutes[event['time'].to_time.min + minute] += 1
  end
end

sleepiest_minute = sleep_minutes.max_by(&:last).first
puts "#1: #{sleepiest_guard_id * sleepiest_minute}"

guard_sleep_minutes = {}
guard_index.each do |id, guard_events|
  sleep_minutes = Hash.new(0)

  guard_events.each_with_index do |event, i|
    next if ['begins shift', 'wakes up'].include?(event['event'])

    next_event = guard_events[i + 1]
    raise 'event order mismatch' if next_event['event'] != 'wakes up'

    seconds_asleep = next_event['time'].to_time - event['time'].to_time
    minutes_asleep = seconds_asleep.to_i / 60
    minutes_asleep.to_i.times do |minute|
      sleep_minutes[event['time'].to_time.min + minute] += 1
    end
  end

  guard_sleep_minutes[id] = sleep_minutes.max_by(&:last).first unless sleep_minutes.empty?
end

p guard_sleep_minutes

most_consistent_guard = guard_sleep_minutes.max_by(&:last)
p most_consistent_guard
puts "#2: #{most_consistent_guard.reduce(:*)}"

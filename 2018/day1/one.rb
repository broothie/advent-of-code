require 'set'

input = File.read('input.txt')
values = input.split("\n").map(&:to_i)

counter = 0
frequency = 0
frequencies = Set.new
loop do
  frequency += values[counter % values.length]
  if frequencies.include?(frequency)
    puts frequency
    break
  end

  frequencies.add(frequency)
  counter += 1
end

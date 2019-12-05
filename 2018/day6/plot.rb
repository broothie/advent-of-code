require 'set'

input = File.read('input.txt')
lines = input.split("\n")

locations = lines.map { |line| line.split(', ').map(&:to_i) }

min_x = max_x = locations.first.first
min_y = max_y = locations.first.last

locations.each do |location|
  x, y = location
  max_x = x if x > max_x
  max_y = y if y > max_y
  min_x = x if x < min_x
  min_y = y if y < min_y
end

x_range = max_x - min_x
y_range = max_y - min_y

p [x_range, y_range]

normalized = locations.map do |location|
  x, y = location
  [x - min_x, y - min_y]
end

location_set = Set.new(normalized)
builder = StringIO.new
counter = 0
y_range.times do |y|
  x_range.times do |x|
    if location_set.include?([x, y])
      builder << 'O'
      counter += 1
    else
      builder << '-'
    end
  end
  builder << "\n"
end

puts builder.string

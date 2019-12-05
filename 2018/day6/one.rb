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
snaked_locations = Set.new

print_map = ->(location) do
  lx, ly = location
  builder = StringIO.new
  (ly-20 .. ly+20).each do |y|
    (lx-20 .. lx+20).each do |x|
      cl = [x, y]
      if location_set.include?(cl)
        builder << 'O'
      elsif snaked_locations.include?(cl)
        builder << '@'
      else
        builder << '.'
      end
    end
    builder << "\n"
  end

  system 'clear'
  puts builder.string
end

class LoopBreak < StandardError; end

# normalized.each do |location|
#   puts "running: [#{location.join(', ')}]"
#   x, y = location
#
#   ring = 0
#   counter = 0
#   begin
#     loop do
#       ring += 1
#
#       counter += 1
#       counter.times do
#         x += 1
#         raise LoopBreak if location_set.include?([x, y])
#         snaked_locations << [x, y]
#       end
#
#       counter.times do
#         y += 1
#         raise LoopBreak if location_set.include?([x, y])
#         snaked_locations << [x, y]
#       end
#
#       counter += 1
#       counter.times do
#         x -= 1
#         raise LoopBreak if location_set.include?([x, y])
#         snaked_locations << [x, y]
#       end
#
#       counter.times do
#         y -= 1
#         raise LoopBreak if location_set.include?([x, y])
#         snaked_locations << [x, y]
#       end
#     end
#   rescue LoopBreak
#     pair = [x, y]
#   end
#
#   puts "[#{pair.join(', ')}] snake ended at: [#{pair.join(', ')}]"
# end

def distance(a, b)
  a.zip(b).map { |d| d.inject(:-).abs }
end

edges_removed = normalized.reject { |x, y| [0, x_range].include?(x) || [0, y_range].include?(y) }

bad_locations = Set.new
close_counts = Hash.new(0)
x_range.times do |x|
  y_range.times do |y|
    point = [x, y]
    closests = edges_removed.sort_by { |l| distance(l, point) }
    closest = closests.first
    next if closests.select { |l| l == closest }.size > 1
    area_on_edge = [0, x_range].include?(x) || [0, y_range].include?(y)
    bad_locations << closest if area_on_edge
    close_counts[closest] += 1
  end
end

close_counts.delete_if { |key| bad_locations.include?(key) }

p close_counts
p close_counts.max_by(&:last)
p close_counts.map(&:last).inject(&:+) == x_range * y_range

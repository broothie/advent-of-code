
input = File.read('input.txt')
w0, w1 = input.split("\n")
w0_dirs, w1_dirs = w0.split(','), w1.split(',')

grid = Hash.new { |h, k| h[k] = [] }
[w0_dirs, w1_dirs].each_with_index do |dirs, wire|
  x, y = 0, 0
  dirs.each do |dir|
    len = dir.chars.drop(1).join.to_i

    case dir.chars.first
    when 'R' then len.times { x += 1; grid[[x, y]] << wire }
    when 'L' then len.times { x -= 1; grid[[x, y]] << wire }
    when 'D' then len.times { y += 1; grid[[x, y]] << wire }
    when 'U' then len.times { y -= 1; grid[[x, y]] << wire }
    end
  end
end

intersections = []
grid.each do |pos, wires|
  intersections << pos if wires.uniq.size > 1
end

puts intersections.map { |pos| pos.map(&:abs).sum }.sort.first

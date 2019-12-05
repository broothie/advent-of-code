
input = File.read('input.txt')
w0, w1 = input.split("\n")
w0_dirs, w1_dirs = w0.split(','), w1.split(',')

grid_steps = Hash.new { |h, k| h[k] = [] }
grid = Hash.new { |h, k| h[k] = [] }
[w0_dirs, w1_dirs].each_with_index do |dirs, wire|
  steps = 0
  x, y = 0, 0
  dirs.each_with_index do |dir|
    len = dir.chars.drop(1).join.to_i

    case dir.chars.first
    when 'R' then len.times { x += 1; grid[[x, y]] << wire; steps += 1; grid_steps[[x, y, wire]] << steps }
    when 'L' then len.times { x -= 1; grid[[x, y]] << wire; steps += 1; grid_steps[[x, y, wire]] << steps }
    when 'D' then len.times { y += 1; grid[[x, y]] << wire; steps += 1; grid_steps[[x, y, wire]] << steps }
    when 'U' then len.times { y -= 1; grid[[x, y]] << wire; steps += 1; grid_steps[[x, y, wire]] << steps }
    end
  end
end

intersections = []
grid.each do |pos, wires|
  intersections << pos if wires.uniq.size > 1
end

lengths = intersections.map do |pos|
  l0 = grid_steps[pos + [0]].min
  l1 = grid_steps[pos + [1]].min
  l0 + l1
end

puts lengths.min

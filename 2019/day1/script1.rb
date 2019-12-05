
lines = File.read('input.txt').split("\n")

masses = lines.map(&:to_i)
total_mass = masses.map { |mass| mass / 3 - 2 }.sum
puts total_mass

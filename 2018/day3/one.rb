
claim_re = /#(?<id>\d+) @ (?<x>\d+),(?<y>\d+): (?<w>\d+)x(?<h>\d+)/

input = File.read('input.txt')
# input = <<INPUT
# #1 @ 1,3: 4x4
# #2 @ 3,1: 4x4
# #3 @ 5,5: 2x2
# INPUT

raw_claims = input.split("\n")

claims = raw_claims.map do |raw_claim|
  match = claim_re.match(raw_claim)
  match.names.zip(match.captures).to_h
end

matrix = []
claims.each do |claim|
  claim['h'].to_i.times do |i|
    claim['w'].to_i.times do |j|
      x = claim['x'].to_i + j
      y = claim['y'].to_i + i

      matrix[y] ||= []
      matrix[y][x] ||= 0
      matrix[y][x] += 1
    end
  end
end

puts "1: #{matrix.flatten.count { |e| e.to_i > 1 }}"

claims.each do |claim|
  overlap = false
  claim['h'].to_i.times do |i|
    claim['w'].to_i.times do |j|
      x = claim['x'].to_i + j
      y = claim['y'].to_i + i

      count = matrix[y][x]
      overlap = true if count > 1
    end
  end

  unless overlap
    puts "2: #{claim['id']}"
    exit
  end
end

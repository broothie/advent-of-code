
input = (152085..670283)

def valid?(num)
  num = num.to_s
  return false unless num.size == 6
  return false unless num.chars.each_cons(2).all? { |pair| pair.first <= pair.last }
  return false unless num.chars.chunk { |e| e }.map(&:last).any? { |chunk| chunk.size == 2 }

  true
end

counter = 0
input.each do |num|
  counter += 1 if valid?(num)
end

puts counter

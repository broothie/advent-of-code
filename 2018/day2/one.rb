require 'set'

input = File.read('input.txt')
codes = input.split("\n")

def count_letters(code)
  code.chars.each_with_object(Hash.new(0)) do |char, memo|
    memo[char] += 1
  end
end

def count_counts(count_hash)
  set = Set.new(count_hash.values)

  two = set.include?(2) ? 1 : 0
  three = set.include?(3) ? 1 : 0

  return two, three
end

def code_counts(code)
  count_counts(count_letters(code))
end

two_counter = three_counter = 0
codes.each do |code|
  two, three = code_counts(code)

  two_counter += two
  three_counter += three
end

puts two_counter * three_counter

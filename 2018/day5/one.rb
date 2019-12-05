
input = File.read('input.txt')
puts input.size

def react(string)
  reacted = false
  chars = string.chars
  (0...chars.size).each do |i|
    next if i.zero?

    char = chars[i]
    prev_char = chars[i - 1]
    next unless prev_char

    if react?(prev_char, char)
      reacted = true
      chars[i] = nil
      chars[i - 1] = nil
    end
  end

  [chars.compact.join, reacted]
end

def react?(a, b)
  a_case = a == a.downcase ? 'lower' : 'upper'
  b_case = b == b.downcase ? 'lower' : 'upper'

  (a_case != b_case) && (a.downcase == b.downcase)
end

def fully_react(string)
  loop do
    string, reacted = react(string)
    break unless reacted
  end

  string
end

def remove_char(string, char)
  chars = string.chars
  (0...chars.size).each do |i|
    curr_char = chars[i]
    chars[i] = nil if [curr_char, curr_char.swapcase].include?(char)
  end

  chars.compact.join
end

# require 'byebug'; byebug
count_hash = {}
('a'..'z').each do |char|
  filtered_input = remove_char(input, char)
  count_hash[char] = fully_react(filtered_input).size
  puts "#{char}: #{count_hash[char]}"
end

p count_hash.min_by(&:last)

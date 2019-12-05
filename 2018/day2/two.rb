
input = File.read('input.txt')
codes = input.split("\n")

def compare_codes(code1, code2)
  count = 0
  (0...code1.length).each do |i|
    char1, char2 = code1[i], code2[i]
    count += 1 if char1 != char2
  end

  count
end

codes.each do |code1|
  codes.each do |code2|
    next if code1 == code2

    count = compare_codes(code1, code2)
    if count == 1
      diff = code1.chars - code2.chars
      puts (code1.chars - diff).join
      exit
    end
  end
end

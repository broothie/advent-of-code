
nums = File.read('input.txt').split(',')
intcode = nums.map(&:to_i)

# before running the program, replace position 1 with the value 12 and replace position 2 with the value 2
intcode[1], intcode[2] = 12, 2

cursor = 0
loop do
  opcode = intcode[cursor]

  lp, rp, tp = intcode[cursor + 1], intcode[cursor + 2], intcode[cursor + 3]
  lv, rv = intcode[lp], intcode[rp]

  case opcode
  when 1 then intcode[tp] = lv + rv
  when 2 then intcode[tp] = lv * rv
  when 99 then break
  end

  cursor += 4
end

puts intcode.first

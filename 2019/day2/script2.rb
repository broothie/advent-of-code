
nums = File.read('input.txt').split(',')
raw_intcode = nums.map(&:to_i)

0.upto(99).each do |noun|
  0.upto(99).each do |verb|
    intcode = raw_intcode.dup
    intcode[1], intcode[2] = noun, verb

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

    if intcode.first == 19690720
      puts noun
      puts verb
      exit
    end
  end
end

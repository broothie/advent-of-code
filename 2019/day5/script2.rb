
raw = File.read('input.txt')
intcode = raw.split(',').map(&:to_i)
inputs = [5]

def value(intcode, mode, arg)
  mode == 1 ? arg : intcode[arg]
end

pc = 0
loop do
  instruction = intcode[pc].to_s
  if instruction.length == 1
    opcode = instruction.to_i
  else
    opcode = instruction.chars.drop(instruction.length - 2).join.to_i
    modes = instruction.chars.take(instruction.length - 2).reverse.map(&:to_i)
  end

  case opcode
  when 1
    la, ra, ta = intcode[pc+1 .. pc+3]
    lm, rm, tm = modes
    intcode[ta] = value(intcode, lm, la) + value(intcode, rm, ra)
    pc += 4

  when 2
    la, ra, ta = intcode[pc+1 .. pc+3]
    lm, rm, tm = modes
    intcode[ta] = value(intcode, lm, la) * value(intcode, rm, ra)
    pc += 4

  when 3
    input = inputs.pop
    a = intcode[pc + 1]
    intcode[a] = input
    pc += 2

  when 4
    a = intcode[pc + 1]
    puts intcode[a]
    pc += 2

  when 5
    ca, ja = intcode[pc+1], intcode[pc+2]
    cm, jm = modes
    if value(intcode, cm, ca) != 0
      pc = value(intcode, jm, ja)
    else
      pc += 3
    end

  when 6
    ca, ja = intcode[pc+1], intcode[pc+2]
    cm, jm = modes
    if value(intcode, cm, ca) == 0
      pc = value(intcode, jm, ja)
    else
      pc += 3
    end

  when 7
    la, ra, ta = intcode[pc+1 .. pc+3]
    lm, rm = modes
    intcode[ta] = value(intcode, lm, la) < value(intcode, rm, ra) ? 1 : 0
    pc += 4

  when 8
    la, ra, ta = intcode[pc+1 .. pc+3]
    lm, rm = modes
    intcode[ta] = value(intcode, lm, la) == value(intcode, rm, ra) ? 1 : 0
    pc += 4

  else
    break

  end
end

File.open(ARGV[0]).each_line do |line|
  s, b, c = line.split.map(&:to_i), -1, 1
  for i in s
    if i == b
      c += 1
    else
      if b >= 0
        print c.to_s + " " + b.to_s + " "
        c = 1
      end
      b = i
    end
  end
  puts c.to_s + " " + b.to_s
end

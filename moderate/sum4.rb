def numzero(s, c, z)
  return z == 0 ? 1 : 0 if c == 0
  return 0 if s.length < c or z > 0
  numzero(s[1..-1], c-1, z+s[0]) + numzero(s[1..-1], c, z)
end

File.open(ARGV[0]).each_line do |line|
  puts numzero(line.chomp.split(',').map(&:to_i).sort, 4, 0)
end

require 'prime'
File.open(ARGV[0]).each_line do |line|
  s, c = line.split(",").map(&:to_i), 0
  (s[0]..s[1]).each do |i|
    c += 1 if Prime.prime?(i)
  end
  puts c
end

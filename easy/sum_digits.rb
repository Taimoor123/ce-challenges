File.open(ARGV[0]).each_line do |line|
  puts line.split('').map(&:to_i).inject{|sum,x| sum + x}
end

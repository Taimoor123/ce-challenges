class String
  def rotate!
    self << self[0]
    self[0] = ""
  end
end

File.open(ARGV[0]).each_line do |line|
  s, r = line.chomp.split(','), false
  s[0].length.times do
    if s[0] == s[1]
      r = true
      break
    else
      s[1].rotate!
    end
  end
  puts r ? "True" : "False"
end

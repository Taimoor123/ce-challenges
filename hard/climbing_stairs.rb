File.open(ARGV[0]).each_line do |line|
  if line !~ /^$/
    n = line.to_i
    if n < 4
      puts n
    else
      r, p = 1, 1
      r, p, n = r+p, r, n-1 while n > 1
      puts r
    end
  end
end

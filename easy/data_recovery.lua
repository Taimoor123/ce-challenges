for line in io.lines(arg[1]) do
  local a, r = {}, {}
  for i in line:gmatch("[^ ;]+") do a[#a+1] = i end
  local h = math.ceil(#a/2)
  for i = 1, h-1 do
    r[tonumber(a[h+i])] = a[i]
  end
  for i = 1, h do
    if r[i] == nil then r[i] = a[h] break end
  end
  for i = 1, #r do
    if i > 1 then io.write(" ") end
    io.write(r[i])
  end
  print()
end
def robot(f, x, y)
  return 1 if x == 3 and y == 3
  ret = 0
  if x > 0 and f & (1 << (x-1 + 4*y)) == 0
    g = f | (1 << (x-1 + 4*y))
    ret += robot(g, x-1, y)
  end
  if y > 0 and f & (1 << (x + 4*(y-1))) == 0
    g = f | (1 << (x + 4*(y-1)))
    ret += robot(g, x, y-1)
  end
  if x < 3 and f & (1 << (x+1 + 4*y)) == 0
    g = f | (1 << (x+1 + 4*y))
    ret += robot(g, x+1, y)
  end
  if y < 3 and f & (1 << (x + 4*(y+1))) == 0
    g = f | (1 << (x + 4*(y+1)))
    ret += robot(g, x, y+1)
  end
  ret
end

puts robot(1, 0, 0)

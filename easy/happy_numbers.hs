import System.Environment (getArgs)

happy    :: Int -> Int -> Int
happy x y | x == 0    = y
          | otherwise = happy (div x 10) (y + z*z)
          where z = mod x 10

testHappy    :: Int -> Int -> String
testHappy i j | i >= 127 || j == 0 = "0"
              | j == 1             = "1"
              | otherwise          = testHappy (succ i) (happy j 0)

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr . unlines . map (testHappy 0 . read) $ lines input
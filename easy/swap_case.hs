import System.Environment (getArgs)
import Data.Char (isLower, isUpper, toLower, toUpper)

swapCase   :: Char -> Char
swapCase c | isLower c = toUpper c
           | isUpper c = toLower c
           | otherwise = c

main = do
    [inpFile] <- getArgs
    input <- readFile inpFile
    putStr $ map swapCase input
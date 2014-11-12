import System.Environment (getArgs)
import System.Posix.Files (fileSize, getFileStatus)

main :: IO ()
main = do
    [inpFile] <- getArgs
    status <- getFileStatus inpFile
    putStrLn . show . toInteger . fileSize $ status

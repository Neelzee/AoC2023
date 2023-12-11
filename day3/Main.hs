module Main where
import Text.Read (readMaybe)


--          Obst x   y
data Obst = Obst Int Int
  deriving Show

--          Digs num x   y
data Digs = Digs Int Int Int
  deriving Show

-- Find all obst in a line
-- Line -> x -> y -> [Obst]
findObst :: String -> Int -> Int -> [Obst]
findObst [] _ _ = []
findObst (c:cs) x y
  | c == '.' = findObst cs (x + 1) y
  | otherwise = case readMaybe [c] :: Maybe Int of
    Just _ -> findObst cs (x + 1) y
    Nothing -> Obst x y : findObst cs (x + 1) y

findDigs :: String -> Int -> Int -> Int -> [Digs]
findDigs [] _ _ _ = []
findDigs (c:cs) x y n = case readMaybe [c] :: Maybe Int of
  Just nm -> findDigs cs (x + 1) y (add n nm)
  Nothing -> if n == 0
    then
      findDigs cs (x + 1) y 0
    else
      Digs n (x - 1) y : findDigs cs (x + 1) y 0
  where
    add :: Int -> Int -> Int
    add a b = a * 10 + b


findAllDigs :: [String] -> [Digs]
findAllDigs = fad 0
  where
    fad :: Int -> [String] -> [Digs]
    fad _ [] = []
    fad y (c:cs) = findDigs c 0 y 0 ++ fad (y + 1) cs

findAllObst :: [String] -> [Obst]
findAllObst = fao 0
  where
    fao :: Int -> [String] -> [Obst]
    fao _ [] = []
    fao y (c:cs) = findObst c 0 y ++ fao (y + 1) cs

numDigits :: Int -> Int
numDigits = length . show

isNextTo :: Digs -> Obst -> Bool
isNextTo (Digs num x1 y1) (Obst x2 y2) =
    horizontalAdjacent || verticalAdjacent || diagonalAdjacent
  where
    numD = numDigits num
    xEnd = x1 + numD - 1 -- The x-coordinate of the last digit in Digs
    horizontalAdjacent = any (\x -> x == x2 && abs (y1 - y2) == 1) [x1..xEnd]
    verticalAdjacent = any (\x -> abs (x - x2) == 1 && y1 == y2) [x1..xEnd]
    diagonalAdjacent = any (\x -> abs (x - x2) == 1 && abs (y1 - y2) == 1) [x1..xEnd]



isNextToAny :: Digs -> [Obst] -> Bool
isNextToAny d = any (isNextTo d)

findAdjacentDigs :: [String] -> [Digs]
findAdjacentDigs xs = calc xs 0
  where
    calc [] _ = []
    calc [_] _ = []
    calc (l1:l2:ls) y =
      let obs = findObst l1 0 y ++ findObst l2 0 (y + 1)
          digs = findDigs l1 0 0 y ++ findDigs l2 0 0 (y + 1)
      in filter (`isNextToAny` obs) digs ++ calc (l2:ls) (y + 1)


main :: IO ()
main = do
  c <- readFile "data"
  print $ findAdjacentDigs (lines c)


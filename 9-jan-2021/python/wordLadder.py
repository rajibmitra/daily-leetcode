import collections

class Solution(object):
    def ladderLength(self, beginWord, endWord, wordList):

        if beginWord in wordList:
            wordList.remove(beginWord)
        
        if endWord not in wordList:
            return 0

        wordList.append(endWord)
        q = collections.deque([[beginWord, 1]])
        leng = len(beginWord)
        
        while q:
            word, length = q.popleft()
            if word == endWord:
                return length
            wordListCopy = wordList.copy()
            for n in wordListCopy:
                if sum(n[i]!=word[i] for i in range(leng))==1:
                    wordList.remove(n)
                    q.append([n,length+1])
        return 0
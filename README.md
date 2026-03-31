# AlgoWithMe
## Purpose
Lucas: I keep going to these different sites leetcode, algoexpert.io, hackerrank and the like and they all forced me to interact with their website in clunky and sometimes even fustrating UIs. So I thought "what if you just used normal developer tools locally and used real development lifecycle processes to learn the data structures & algorithms that help us level up our ability to engineer: test locally, commit, push, PR, GitHub workflows?". I believe it would look like AlgoWithMe :)

By using this you are getting real experience so use it as a chance find new tools to use for editing code, working with git and the like. The whole point of doing this instead of the web-based is that you're using a "real" process of testing your skills and learning.   

I will also add go ahead and cheat, figure out a way around the encrypted tests but it doesn't really help you or anyone else. You could also look at other peoples PR for the problem you're working on but what is the point... Why do this at all if you're just cheating, it'd be a waste of time.

Please don't ruin other's experience either, but if others ask for help try to give it to them without giving a solution. This whole thing is for learning.

## How it works
- [Fork this repo into your own github account](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/fork-a-repo)
- [Create a feature branch](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-and-deleting-branches-within-your-repository) (ex: feature/twosumsolution)
- Make changes to the code and [commit them](https://github.com/git-guides/git-commit)
  - Run tests locally, look at .github/workflows/{go_tests.yml,go_more_tests.yml} to see how it is being tested. Always read the code. It's a good way to get familiar with things.
- [Push](https://github.com/git-guides/git-push) changes to your forked repo feature branch
  - This will run the first level of tests
- [Create a Pull Request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request)
  - After creating the Pull Request ask @catdevman to add a label for checking the second level of tests.
- If the tests succeed on the second level then you completed that question CONGRATS!
- If the tests fail
  - Change the code, commit, push, and asked @catdevman on the Pull Request to rerun the tests (yeah I know this isn't great but only way I can keep the second level of tests "secret") 
- Have fun and feel free to ask @catdevman any questions whether about AlgoWithMe or DSA I'll do whatever I can to help

## Resources for learning DSA (no sponsors)
- https://frontendmasters.com/courses/algorithms/
- https://frontendmasters.com/courses/advanced-algorithms/
- https://www.youtube.com/watch?v=AGGqPhBvZTg
- https://www.enjoyalgorithms.com/blog/step-by-step-guidance-to-master-data-structure-and-algorithms-for-coding-interview


I'll probably use LLMs to generate more questions, it would take a lot of effort and time I don't have to do it any other way. However, if you want to contribute feel free to ask.


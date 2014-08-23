##Lab 2 instructions

*Please consult the general workflow and handin instructions below; they are
somewhat different from those of lab 1.* For lab 2, you should edit the file,
ANSWERS.md as described within the file. All other source files must be
submitted as part of the pull request.

##Workflow and Handin Instructions:

For every lab, you are required to write a brief report in the form of answers
to questions in a text file formated as markdown. The labs also requires that
you submit your code. Submission of these deliverables shall be done through
the git repository. Below we describe the workflow that we expect you to follow
for lab 2.

1. Click on the *Fork* button (upper right-hand corner of the page).
2. Follow the on-screen instructions.
3. Once you have a fork of the repository connected to your own user, you will use the `go get` command to clone *this repository* (not your own fork) into your go environment. Here is how to do it:
  - On the command line enter:
  		`go get github.com/uis-dat320-fall2014/lab2`.
    This will clone the original lab 2 git repo (not your fork of it.) This is
    important because it means that you don't need to change the import path in
    the source files to use your own forked repository's path. That is, when
    you make a pull request to submit your handin, you don't have to change
    this back to the original import path.
4. Next, run the following command: 
		`git remote add lab2 https://github.com/username/lab2.git`
	where `username` should be replaced with your usersname.
5. This command adds your own `lab2` fork as a remote, which means that once you've modified some files and committed the changes you can now run:
		`git push lab2`
	to have them pushed up to your lab2 fork on github.

###Ready to submit?
1. When you are finished with all the exercises in the lab, and wish to submit, then first make sure you have pushed your changes to github using: `git push lab2`.
2. Next, go to original lab project (here) and click on the *Pull Requests* link in the right-hand margin menu, and on the page that comes up next, click on the green button titled *New Pull Request*.
3. The next page shows (at the top): *We’re showing branches in this repository, but you can also compare across forks*. Click the *compare across forks* link. In the dropdown menu that appears, select your own fork.
4. You will now see your own changes compared to any original files. On this page, select the green *Create pull request* button.
5. In the title field, write only the following: `username labX submission`, where you replace `username` with your github username and `X` with the lab number. If there are any issue you want us to pay attention to, please use the comment field. Otherwise, leave it blank.
6. Finally, click the green *Create pull request* button.
7. Now, we will be able to review your answers.


# commandgo

Going commando on your boilerplate

I think we all know that Go's approach to passing errors as return values is here to stay,
and I can accept the rather perfuse appearance of errors in my code, despite what I think of it.

However, I find it too distasteful in my tests, where I just wan't to see what's being tested, not
how I managed to set up and consume a method, struct, or command line command

Further, I think it's really important to show what is being changed about the parameters to the SUT
right inside the declaration of a Ginkgo Context. It takes effort to achieve this, and some planning.

In this library I specifically set out to:

  * replace common error boilerplate
    
    ```
    _, err := call()
    if (err != nil) {
      // fail test in some way
    }
    ```
    
    with code that assumes you always want setup to fail (or handle all setup errors orthogonally).
    
    **Note:** this means a lot of what you see in this lib will effectively be a decorator that removes the `error` from
    function return types of standard library functions.
    
    This code from `fixture.go` is an example. It also makes the method fluent, for a little more elegance in your tests.
    
    ```
    func (fixture *Fixture) Write(bytes []byte) *Fixture {
    	_, writeError := fixture.file.Write(bytes)
    
    	if (writeError != nil) {
    		fixture.ErrorHandler(writeError)
    	}
    
    	return fixture
    }
    ```
      
  
  * make it easier to set up SUTs with incremental context changes in the BeforeEach block.

## Fixture

If you're dealing with input files and you want to declare them inline (versus as a file on disk 
committed with your code) then this is the struct for you.

```
  BeforeEach(func() {
    fixtureText :=
  `---
  attribute: value`
    
    fixture := NewFixture("my-great-file").
      Write([]byte(fixtureText)).
      Close()
  
    executable = NewExecutableContext("path/to/executable", "-f", fixture.Name())
  }
```

or use [heredoc](https://github.com/MakeNowJust/heredoc)

```
  import (
    . "github.com/MakeNowJust/heredoc/dot"
  )
  
  BeforeEach(func() {
    fixtureText := D(`
      ---
      attribute: value
    `)
    
    fixture := NewFixture("my-great-file").
      Write([]byte(fixtureText)).
      Close()
  
    executable = NewExecutableContext("path/to/executable", "-f", fixture.Name())
  }
```

## ExecutableContext

Testing executable commands In Go is about building up contexts, so `exec.Command` 
gets replaced with `ExecutableContext.Command`, and ExecutableContext is a context:

```
  var _ = Describe("My Cool Executable, be-cool subcommand", func() {

    var executable *ExecutableContext
  
    BeforeEach(func() {
      executable = NewExecutableContext("path/to/my/executable, "be-cool")
    })
  
    Context("with verbose mode on", func() {
      BeforeEach(func() {
        executable.AddArguments("-v")
      })
  
      Context("Passing a valid cool-factor", func() {
        BeforeEach(func() {
          executable.AddArguments("--cool-factor", "chillaxed")
        })
  
        It("makes things cool", func() {
          command, stdin := executable.PipeCommand()
  
          session, err := Start(command, GinkgoWriter, GinkgoWriter)
          Expect(err).ToNot(HaveOccurred())
  
          stdin.Write([]byte("Fevered feeling, hot hot hot!"))
          stdin.Close()
  
          Eventually(session).Should(Say("Shivery feeling, chillaxed chillaxed chillaxed!"))
          Eventually(session).Should(Exit(0))
        })
      })
  
      Context("Passing an invalid cool-factor", func() {
        BeforeEach(func() {
          executable.AddArguments("--cool-factor", "nerdy")
        })
  
        It("is embarrassing", func() {
          command, stdin := executable.PipeCommand()
  
          session, err := Start(command, GinkgoWriter, GinkgoWriter)
          Expect(err).ToNot(HaveOccurred())
  
          stdin.Write([]byte("Fevered feeling, hot hot hot!"))
          stdin.Close()
  
          Eventually(session).Should(Say("No, dude, you aint cool... But you are excellent :D"))
          Eventually(session).Should(Exit(0))
        })
      })
    })
  })
```

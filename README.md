# commandgo

Going commando on your boilerplate

Testing executable commands In Go is about building up contexts, so `exec.Command` 
gets replaced with `Runner.Command`, and Runner is a context:

```
var _ = Describe("My Cool Executable, be-cool subcommand", func() {

	var runner *Runner

	BeforeEach(func() {
		runner = NewRunner("path/to/my/executable, "be-cool")
	})

  Context("with verbose mode on", func() {
    BeforeEach(func() {
      runner.AddArguments("-v")
    })

    Context("Passing a valid cool-factor", func() {
      BeforeEach(func() {
        runner.AddArguments("--cool-factor", "chillaxed")
      })

      It("makes things cool", func() {
        command, stdin := runner.PipeCommand()

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
        runner.AddArguments("--cool-factor", "nerdy")
      })

      It("is embarrassing", func() {
        command, stdin := runner.PipeCommand()

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

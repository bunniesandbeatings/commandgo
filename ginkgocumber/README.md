```
  Given("I am a developer of mixed tastes", func() {
    developer := Developer(Tastes: tastes.Mixed)
    
    When("I use Given-When-Then", func() {
      developer.Use("Given When and Then")
     
      Then("I am productive", func() {
         Expect(developer.outputLevel).To(Equal(outputLevels.Productive))
      })
    })
  })
```

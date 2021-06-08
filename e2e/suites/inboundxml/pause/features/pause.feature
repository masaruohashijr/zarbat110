Feature: pause    
  In order to pause a sequence of phrases
  As an end user
  I want to call to a Number and say some words and phrases pausing some seconds between each sentence.

  Scenario: Pause a sequence of sentences

    Given I have my "Account"
    When I make a call to "+999"
    And say the sentence "I see trees of green"
    And pause 5 seconds
    And say the sentence "I see skies of blue"
    Then I time this call with more than 5 seconds
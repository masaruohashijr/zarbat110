Feature: redirect    
  In order to redirect a call
  As an end user
  I want to make a call to a Number and listen to a message from the redirected Url.

  Scenario: Redirect to receive a Speech

    Given I have my "Account"
    And I set "+888" to gather speech with 5 seconds of timeout
    And I set "+999" a "woman" to say "You will be redirected."
    And I set "+999" to redirect to "redirect_url"
    And I set "redirect_url" a "woman" to say "You was redirected."
    When I try to make a call from "+888" to "+999"    
    Then I get a "You will be redirected. You was redirected." response in plain text
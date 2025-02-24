Feature: Get collection of boardgames

  Scenario: Get a collection for a valid user
    Given the API is initialized with a valid base URL and HTTP client
    When I request the collection for user "richardwooding"
    Then I should receive a collection of boardgames
    And the collection should contain boardgames

  Scenario: Get a collection for a non-existent user
    Given the API is initialized with a valid base URL and HTTP client
    When I request the collection for user "nonExistentUser"
    Then I should receive an error message
    And the error message should indicate that the user was not found

  Scenario: Get a collection with an empty username
    Given the API is initialized with a valid base URL and HTTP client
    When I request the collection with an empty username
    Then I should receive an error message
    And the error message should indicate that the username is invalid

Feature: Get geeklist by ID

  Scenario: Get a specific geeklist by ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the geeklist with ID "11205"
    Then I should receive a single geeklist
    And the geeklist should have the ID "11205"

#  Scenario: Get a non-existent geeklist by ID
#    Given the API is initialized with a valid base URL and HTTP client
#    When I request the geeklist with ID "00000"
#    Then I should receive an error message
#    And the error message should indicate that the geeklist was not found

  Scenario: Get geeklist with an empty ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the geeklist with an empty ID
    Then I should receive an error message
    And the error message should indicate that the ID is invalid

  Scenario: Get geeklist with comments included
    Given the API is initialized with a valid base URL and HTTP client
    When I request the geeklist with ID "11205" with comments included
    Then I should receive a single geeklist
    And the geeklist should have the ID "11205"
    And the geeklist should have comments

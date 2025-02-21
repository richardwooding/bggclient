Feature: Get boardgames by ID

  Scenario: Get a specific boardgame by ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "12345"
    Then I should receive a boardgame
    And the boardgame should have the ID "12345"

  Scenario: Get multiple boardgames by IDs
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgames with IDs
      | 12345 |
      | 67890 |
    Then I should receive a list of boardgames
    And the list should contain a boardgame with the IDs
      | 12345 |
      | 67890 |

  Scenario: Get a non-existent boardgame by ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with ID "00000"
    Then I should receive an error message
    And the error message should indicate that the boardgame was not found

  Scenario: Get boardgames with an empty ID
    Given the API is initialized with a valid base URL and HTTP client
    When I request the boardgame with an empty ID
    Then I should receive an error message
    And the error message should indicate that the ID is invalid
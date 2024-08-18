package main

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

type PasswordRequest struct {
	Password string `json:"init_password" binding:"required"`
}

type StepResponse struct {
	Steps int `json:"num_of_steps"`
}

// checkPasswordHandler processes the password and returns its strength
func checkPasswordHandler(c *gin.Context) {
	var req PasswordRequest

	// Bind the JSON from the request body to the PasswordRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	// Call your strongPasswordChecker function with the password
	result := strongPasswordChecker(req.Password)
	responseSteps := StepResponse{Steps: result}

	// Return the result as a JSON response
	c.JSON(http.StatusOK, responseSteps)

	// 	numOfSteps := strongPasswordChecker(requestPassword.InitialPassword)
	// 	c.JSON(http.StatusCreated, responseSteps)
}

func strongPasswordChecker(s string) int {
	n := len(s)

	// Variables to check the presence of lowercase, uppercase, and digits
	hasLower := 0
	hasUpper := 0
	hasDigit := 0

	// Replace counter for repeated characters
	replace := 0

	// Count of groups of repeating characters
	oneSeq := 0
	twoSeq := 0

	for i := 0; i < n; {
		if unicode.IsLower(rune(s[i])) {
			hasLower = 1
		}
		if unicode.IsUpper(rune(s[i])) {
			hasUpper = 1
		}
		if unicode.IsDigit(rune(s[i])) {
			hasDigit = 1
		}

		// Check for sequences of repeating characters
		j := i
		for j < n && s[i] == s[j] {
			j++
		}

		seqLen := j - i
		if seqLen >= 3 {
			replace += seqLen / 3
			if seqLen%3 == 0 {
				oneSeq++
			} else if seqLen%3 == 1 {
				twoSeq++
			}
		}

		i = j
	}

	missingTypes := 3 - (hasLower + hasUpper + hasDigit)

	if n < 6 {
		return max(missingTypes, 6-n)
	} else if n <= 20 {
		return max(missingTypes, replace)
	} else {
		delete := n - 20
		replace -= min(delete, oneSeq*1) / 1
		replace -= min(max(delete-oneSeq, 0), twoSeq*2) / 2
		replace -= max(delete-oneSeq-2*twoSeq, 0) / 3
		return delete + max(missingTypes, replace)
	}
}

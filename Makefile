DIR := 2024
DAY ?= 1

BASIC_FILES := input.txt test.txt part1.go part1_test.go part2.go part2_test.go

all: create_files

new:
	@mkdir -p $(DIR)/day$(DAY)
	@for file in $(BASIC_FILES); do \
		touch $(DIR)/day$(DAY)/$$file; \
	done
	@echo "Created new day $(DAY) files"

.PHONY: all create_files

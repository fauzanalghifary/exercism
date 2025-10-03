class Grep
  def self.grep(pattern, flags, files)
    show_line_numbers = flags.include?('-n')
    case_insensitive = flags.include?('-i')
    filename_only = flags.include?('-l')
    exact_match = flags.include?('-x')
    inverted = flags.include?('-v')
    multiple_files = files.length > 1

    results = []

    files.each do |filename|
      file_lines = File.readlines(filename)

      file_lines.each_with_index do |line, index|
        line = line.chomp
        match = if exact_match
                  if case_insensitive
                    line.downcase == pattern.downcase
                  else
                    line == pattern
                  end
                else
                  if case_insensitive
                    line.downcase.include?(pattern.downcase)
                  else
                    line.include?(pattern)
                  end
                end

        match = !match if inverted

        if match
          if filename_only
            results << filename
            break
          else
            result_line = ""
            result_line += "#{filename}:" if multiple_files
            result_line += "#{index + 1}:" if show_line_numbers
            result_line += line
            results << result_line
          end
        end
      end
    end

    results.uniq.join("\n")
  end
end

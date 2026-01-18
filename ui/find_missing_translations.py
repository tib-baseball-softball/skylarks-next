import sys
import re

def find_missing(filename):
    with open(filename, 'r', encoding='utf-8') as f:
        content = f.read()

    # Simple regex to find msgid and msgstr pairs
    # This might not handle all edge cases but should be good enough for standard .po files
    entries = re.findall(r'msgid\s+("(?:[^"\\]|\\.)*"(?:\s*"(?:[^"\\]|\\.)*")*)\s+msgstr\s+("(?:[^"\\]|\\.)*"(?:\s*"(?:[^"\\]|\\.)*")*)', content, re.DOTALL)

    missing = []
    for msgid, msgstr in entries:
        # msgid and msgstr are quoted strings, possibly multiline
        # Strip quotes and joining them
        id_str = "".join(re.findall(r'"((?:[^"\\]|\\.)*)"', msgid))
        str_str = "".join(re.findall(r'"((?:[^"\\]|\\.)*)"', msgstr))

        if id_str and not str_str:
            missing.append(id_str)

    return missing

if __name__ == "__main__":
    for arg in sys.argv[1:]:
        missing = find_missing(arg)
        if missing:
            print(f"File: {arg}")
            for m in missing:
                print(f"  Missing: {m}")

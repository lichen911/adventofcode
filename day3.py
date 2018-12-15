input_file = 'day3_input.txt'

# Claim object to parse the text claim string and model all parts of claim
class Claim(object):
    def __init__(self, claim):
        self.claim_id = int(claim.split('@')[0][1:].strip())
        self.left_margin = int(claim.split('@')[1].split(':')[0].split(',')[0].strip())
        self.top_margin = int(claim.split('@')[1].split(':')[0].split(',')[1].strip())
        self.width = int(claim.split('@')[1].split(':')[1].strip().split('x')[0])
        self.height = int(claim.split('@')[1].split(':')[1].strip().split('x')[1])
        self.no_overlaps = True

    def __str__(self):
        return('#{} @ {},{}: {}x{}'.format(self.claim_id,
                                          self.left_margin,
                                          self.top_margin,
                                          self.width,
                                          self.height))


with open(input_file, 'r') as fd:
    text_claim_list = fd.read().splitlines()

print(text_claim_list)

# populate list of Claim objects
claim_list = []
for text_claim in text_claim_list:
    new_claim = Claim(text_claim)
    claim_list.append(new_claim)

# create 3 dimensional list to model the sheet of fabric.
# the third dimension is a list of Claim objects in that square inch
fabric_sheet = [[[] for _ in range(1000)] for _ in range(1000)]
# populate fabric sheet with all the Claims
for claim in claim_list:
    for col in range(claim.left_margin, claim.width + claim.left_margin):
        for row in range(claim.top_margin, claim.height + claim.top_margin):
            fabric_sheet[col][row].append(claim)

            # set Claim.no_overlaps to False for all claims in a particular square
            # this makes it easy to find the claim with no overlaps
            if len(fabric_sheet[col][row]) >= 2:
                for overlapped_claim in fabric_sheet[col][row]:
                    overlapped_claim.no_overlaps = False

# identify number of fabric squares that have 2 or more overlaps
overlap_count = 0
for col in fabric_sheet:
    for sqr_inch in col:
        if len(sqr_inch) >= 2:
            overlap_count += 1

print(overlap_count)

# find the claim with no overlaps
for claim in claim_list:
    if claim.no_overlaps:
        print(claim)

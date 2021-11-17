# yq2md

yq2md is a simple command line utility to convert a simple yaml file having the
syntax implementation shown under `Example` to a markdown syntax. The output is
spilt to the stdout so you can imagine some creative things starting here

A sample file `sample.yaml` file is also provided to tinker around with.

## Example

yq2md converts

```yaml
- name: "Unit 1: Rotational Dynamics and Oscillatory Motion"
  items:
    - question: |
        (1) **[Example 8.1]** A balance scale consisting of a weightless rod has a mass of 0.1 kg on the right side 0.2 m from a pivot point.

        (a) How far from the pivot point on the left must 0.4 kg be placed so that balance is achieved?
        (b) If the 0.4 kg mass is suddenly removed, what is the instantaneous rotational acceleration of the rod?
        (c) What is the instantaneous tangential acceleration of the 0.1 kg mass when the 0.4 kg mass is removed?
      answer: |
        0.05 m, 49 rad/s$^2$, 9.8 m/s$^2$
    - question: |
        (2) **[Example 8.2]** A large wheel of radius 0.4 m and moment of inertia 1.2 kgm², pivoted at the center is free to rotate without friction. A rope is wound around it and a 2 kg weight is attached to the rope. When the weight has descended 1.5 m from its starting point

        (a) What is it's downward velocity?
        (b) What is the rotational velocity of wheel?
      answer: |
        2.5 m/s, 6.2 rad/s
```

into

```sh
$ yq2md sample.yaml
```

```md
# Unit 1: Rotational Dynamics and Oscillatory Motion
 
(1) **[Example 8.1]** A balance scale consisting of a weightless rod has a mass of 0.1 kg on the right side 0.2 m from a pivot point.
 
(a) How far from the pivot point on the left must 0.4 kg be placed so that balance is achieved?
(b) If the 0.4 kg mass is suddenly removed, what is the instantaneous rotational acceleration of the rod?
(c) What is the instantaneous tangential acceleration of the 0.1 kg mass when the 0.4 kg mass is removed?
 
$\hrulefill$
 
(2) **[Example 8.2]** A large wheel of radius 0.4 m and moment of inertia 1.2 kgm², pivoted at the center is free to rotate without friction. A rope is wound around it and a 2 kg weight is attached to the rope. When the weight has descended 1.5 m from its starting point
 
(a) What is it's downward velocity?
(b) What is the rotational velocity of wheel?
 
$\hrulefill$
```

## Installation

With the go toolchain

```sh
go install github.com/mrsafalpiya/yq2md@latest
```

## Usage

Just RTFM

```sh
$ yq2md --help
```

```
Usage of yq2md:
  -a, --all             Enable this flag to dump all questions into a single root category (disable categorization)
      --html            Produce markdown optimized for HTML conversion
  -n, --numerize        Enable this flag to numerize items within categories
  -r, --randomize       Enable this flag to randomize items within categories
  -s, --show-answer     Enable this flag to answers of the questions too
  -t, --toggle-answer   Enable this flag to make the answer toggle (Available only when --html flag is enabled)
```

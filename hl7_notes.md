# HL7 Parsing Logic (Corepoint)

Below describes all of the cases handled for each HL7 message type, for each
outbound configuration. If it is not listed here, it does not exist--unless
otherwise noted, all other cases are sent "unparsed".

## Outbound MI Configuration

### MI ADT

- MSH.4
  - if "STRIC", copy its value to MSH.5
  - if "VVRMC", do logic on PID/IN1 segments
- value at PID.3 determines what is overwritten to MSH.4
  - "B" = BRN_72
  - "BA" = MSUTH_26
  - "J" = TXSAN_59
  - "K" = SOAK_48
  - "M" = MTR_30, QUAR_82 (2 separate messages sent)
  - "MHHC" = HCM_36
  - "N" = MST_27
  - "V" = NMTH_24, CONV_97, MNACO_106 (3 separate messages sent)
  - "W" = CHLD_06, MTH_03, BRN_72, METHWO_96, MTHDZ_104, MTHLG_105 (6 separate messages sent!)

### MI ORM

- MSH.5
  - if "VVRMC", just format DOB at PID.7
  - if "METH" or "STRIC":
    - derivative is used (OBX-5 is a component, not just text)
    - separate IN1.2 (component), putting IN1.2.2 into IN1.4
    - STRIC only: copy OBR.3 to PID.18

### MI ORU

- Replaces **all** subcomponent delimiters (&) with component (^)
  - except for delimiters field (MSH.2)
- Preliminary reports exluced (OBR.25 = "P")
- MSH.5
  - all cases:
    - copy OBR.6 to OBR.7
    - copy OBR.31 to OBR.43
    - copy "A" to OBR.25 if "Addendum Begins" in any OBX-5 (report)
  - if "STRIC":
    - copy OBR.3 to PID.18
  - all other cases:
    - if MSH.4 = "VVRMC":
      - don't send if OBR.32.1 = "578" (not sure why, I think blocking specific provider)
      - format DOB in PID.7

## Outbound PB Configuration

### PB ORM

- if MSH.4 = "AGFA":
  - copy "SC" to ORC.1
  - copy "CM" to ORC.5
  - copy OBR.7 to OBR.8
  - clear OBR.7

  Additional mapping defined per the below table. If the condition in column 1
  applies, the subsequent columns are assigned.

  | PID.3.6 | PID.3.4 | MSH.4  | PV1.42.1  | PV1.42.2                          |
  |:-------:|:-------:|:------:|:---------:|:---------------------------------:|
  | VVRMC   | VALV    | VVRMC  | Val Verde | Val Verde Regional Medical Center |
  | FRIO    | FRIO    | FRIO   | Frio      | Frio Regional Hospital            |
  | MEDINA  | MEDINA  | MEDINA | Medina    | Medina Regional Hospital          |
  | MHHC    | MHHC    | MHHC   | Methodist | Methodist Hospital Hill Country   |
  | METHWO  | METHWO  | METHWO | Methodist | Methodist Hospital Westover Hills |

- else:
  - PID.3.6
    - if "STRIC":
      - copy PID.3.6 to PV1.3.4
      - if MSH-4 is boutique mammo:
        - copy "BOUT" to PID.3.4
        - copy "STRIC" to PV1.41
      - else:
        - copy "STR" to PID.3.4
        - copy "STRIC" to PV1.41
      - clear PID fields, PV1 (demographics)
      - copy ORC.3 to ORC.2
      - copy OBR.3 to OBR.2
      - copy OBR.27.6 to OBR.5
      - Assign timestamp from OBR.8 to appropriate field based on ORC.5:
        - "CM": copy OBR.8.5 to OBR.8 :woozy-face:
        - "IP": copy OBR.8.4 to OBR.7
        - "SC": copy OBR.8.2 to OBR.7
        - "CA": copy OBR.8.6 to OBR.7
    - else: 
      - copy PID.3.6 to PV1.3.5, MSH.4
      - if MSH.4 matches regex "HOAK|^H$|^BO$|^MC$|MDIC|^DT$|^NC$|^SO$|^NE$|NEHMR|^NW$|TCA|SCHERTZ|WOHILLS|^SS$|^ST$|^T$|^WC$|ASCSO|KHIC|WICOH|LMIC|BOWIC":
        - copy "STR" to PID.3.4
        - copy "STRIC" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 matches regex "^B$|^C$|COCMHL|COCMHL|COCMSO|COCMT|COCMTA|COCSN|COCVO|COCWM|^M$|METHHS|METHMT|METHSO|METHTX|METHNE|^MQ$|MSTH|^WH$|^KR$|^HL$|^NC$|^DZ$|CB|NEWBR|NB|METHLM":
        - copy "METH" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 "BMCAH|BMCAR|BMC|BMCNE|BMCSO":
        - copy "BOUT" to PID.3.4
        - copy "STRIC" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 = "METHWO":
        - copy "METHWO" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 = "MHHC":
        - copy "MHHC" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 = "MEDINA":
        - copy "MEDINA" to PID.3.4
        - copy "Medina" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 = "FRIO":
        - copy "FRIO" to PID.3.4
        - copy "Frio" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 matches regex "ASCMC|ASCBO":
        - copy "METH" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - copy PID.4 to PID.3 (seems contradictory)
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 matches regex "^CV$|CONV_97":
        - copy "CV" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 matches regex "^LT$|MTHLG_105":
        - copy "METH" to PID.3.4
        - copy "Methodist" to PV1.42.1
        - copy "LT" to MSH-4
        - map fullname to PV1-42.2 from PID.3.6
      - if MSH.4 matches regex "VVRMC|VVRMCL18":
        - copy "VALV" to PID.3.4
        - copy "Val Verde" to PV1.42.1
        - map fullname to PV1-42.2 from PID.3.6
      - map patient type to PV1.18 from table (likely incomplete)
      - copy PV1.3.2 to PV1.3.1 if PV1.3.1 is empty
      - copy ORC.3 to ORC.2
      - copy OBR.3 to OBR.2
      - copy OBR.27.6 to OBR.5

### PB MDM (ORU)

**NOTE**: if the following top-level condition is not matched, **the message is not sent**
- if MSH-4 matches regex "METHHS|COCMHL|METHSO|COCMSO|METHMT|COCMT|^MQ$|COCMT|METHTX|COCMTA|MSTH|COCSN|METHNE|COCVO|^B$|^C$|^M$|^WH$|^CV$|BMCH|BMCAH|BMCNE|BMCSO|^BO$|^DT$|^H$|HOAK|TCA|MASH|MC|MDIC|MEDINA|SO|NE|BMCAR|WOHILLS|NEHMR|^ST$|^NW$|ASCMC|ASCSO|^SS$|SCHERTZ|^T$|VVRMCL18|VVRMC|^WC$|LMIC|^LT$|MTHLG_105|Val Verde Regional Medical Ctr|^HL$|^NC$|^DZ$|ASCBO|^KR$|ASCMC|FRIO|CB|NEWBR|KHIC|WICOH|Methodist Hosp Hill Country|NB|METHLM|METHWO|Methodist Hosp Westover Hills|BOWIC|BV":
  - I'm not gonna go through everything. It's a lot of the same stuff in the PB ORM configuration.

## Volta mapping

### MSH (all message types)

| Name                | HL7 Field | Type   | Notes                                                                             |
|:-------------------:|:---------:|:------:|:---------------------------------------------------------------------------------:|
| Sending Application | MSH.3     | string | Majority of ADTs are missing this--why?                                           |
| Sending Facility    | MSH.4     | string | Similarly, blank for many ADTs. Alarming number of ORUs from Royal are also blank |
| Message Type        | MSH.9     | cm_msg |                                                                                   |
| Control ID          | MSH.10    | string |                                                                                   |
| Version             | MSH.12    | string |                                                                                   |

### PID (all message types)

| Name                | HL7 Field | Type   | Notes                                                              |
|
| Patient MRN         | PID.3     | cx     | Why is the "true" site code stored here instead of the MSH.4?      |
| Patient Name        | PID.5     | xpn    |                                                                    |
| DOB                 | PID.7     | ts     |                                                                    |

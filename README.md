
# resource-scheduling
**INTELLIGENT RESOURCE SCHEDULING FOR DISTRIBUTED ORCHESTRATION SYSTEMS USING REINFORCEMENT LEARNING**

### Paper Information
- **Author(s):** Kalesha Khan Pattan
- **Published In:** *****************************
- **Publication Date:** **************
- **ISSN:** *****************
- **DOI:**
- **Impact Factor:** ***********

### Abstract
Distributed orchestration platforms like Kubernetes rely on heuristic-based schedulers that struggle to adapt to dynamic and heterogeneous workloads, leading to inefficiencies 
and performance degradation. This study proposes an intelligent resource scheduling framework using reinforcement learning that continuously learns cluster and workload behavior 
to optimize scheduling decisions. The RL-based scheduler minimizes pod scheduling latency, balances resource utilization, and improves throughput through adaptive, reward-driven 
learning. Experiments conducted on Kubernetes clusters of varying sizes demonstrate consistent performance gains over the default scheduler, even under high contention and node 
failure scenarios. The results highlight reinforcement learning as a robust, scalable, and forward-looking approach for next-generation distributed orchestration systems.

### Key Contributions
- **Intelligent Reinforcement Learning–Based Scheduling Framework:**
  Designed an adaptive resource scheduling framework for distributed orchestration systems that replaces static, heuristic-driven Kubernetes scheduling with reinforcement learning–based decision-making.

  - **Adaptive Exploration–Exploitation Scheduling Strategy:**
  Introduced an epsilon-greedy reinforcement learning policy that dynamically balances exploration and exploitation, enabling continuous learning from scheduling outcomes under changing cluster conditions.

- **Latency-Aware and Load-Balanced Pod Placement:**
  Formulated scheduling as a sequential decision problem that minimizes pod scheduling latency while improving load distribution across heterogeneous cluster nodes.
   
- **Scalable Evaluation Across Multi-Node Clusters:**
  Conducted extensive empirical evaluation on Kubernetes clusters with 3, 5, 7, 9, and 11 nodes, demonstrating consistent and scalable performance gains over the default scheduler.

  ### Relevance & Real-World Impact
- **Consistent Reduction in Scheduling Latency:**
  Achieved sustained latency reductions of approximately 27–37% compared to the baseline scheduler, with increasing benefits observed as cluster size and workload intensity grow.
  
- **Improved Scalability for Large Clusters:**
Demonstrated smoother and more controlled latency growth under scale, addressing a critical limitation of heuristic-based scheduling in large Kubernetes deployments.

- **Scalable Distributed Deployment:**
    Enabled efficient scaling across increasing cluster sizes by dynamically adapting replica strategies to workload intensity and node behavior.

  **Adaptive Resource Utilization Under Dynamic Workloads:**
  Enabled real-time adaptation to workload fluctuations, resource contention, and node variability through reinforcement learning–driven policy updates.
    
- **Robustness in Realistic Cluster Conditions:**
    Showed strong resilience under high-load and failure-prone scenarios, highlighting the suitability of the approach for production-grade cloud-native systems.
    
 - **Academic, Research, and Industry Value:**
  Provides a comprehensive reference framework—covering scheduling models, RL policies, simulations, and evaluation metrics—for researchers, advanced coursework, and practitioners building intelligent orchestration platforms.

### Experimental Results (Summary)

  | Nodes | Baseline Scheduler(ms) | RL Based Scheduler (ms)| Improvment (%)  |
  |-------|------------------------| -----------------------| ----------------|
  | 3     |  85                    | 62                     | 27.06           |
  | 5     |  110                   | 78                     | 29.09           |
  | 7     |  145                   | 100                    | 31.03           |
  | 9     |  190                   | 125                    | 34.21           |
  | 11    |  260                   | 165                    | 36.54           |

### Citation
INTELLIGENT RESOURCE SCHEDULING FOR DISTRIBUTED ORCHESTRATION SYSTEMS USING REINFORCEMENT LEARNING
* Kalesha Khan Pattan
* ******************************************************************** 
* ISSN 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
****************https://www.ijisae.org/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/**** | **Email**: pattankalesha520@gmail.com






